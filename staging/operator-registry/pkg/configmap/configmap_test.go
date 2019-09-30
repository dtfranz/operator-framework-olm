package configmap

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		assertFunc func(t *testing.T, manifestGot *Manifest)
	}{
		{
			name:   "BundleWithCsvAndCrdOnly",
			source: "testdata/bundle.cm.yaml",
			assertFunc: func(t *testing.T, manifestGot *Manifest) {
				assert.NotNil(t, manifestGot.Bundle)
				assert.Nil(t, manifestGot.PackageManifest)

				// We don't populate package, channel and name of a bundle object.
				assert.Empty(t, manifestGot.Bundle.Package)
				assert.Empty(t, manifestGot.Bundle.Channel)
				assert.Empty(t, manifestGot.Bundle.Name)

				csvGot, errGot := manifestGot.Bundle.ClusterServiceVersion()
				assert.NoError(t, errGot)
				assert.NotNil(t, csvGot)

				crdListGot, errGot := manifestGot.Bundle.CustomResourceDefinitions()
				assert.NoError(t, errGot)
				assert.Equal(t, 1, len(crdListGot))
			},
		},

		{
			name:   "BundleWithPackageManifest",
			source: "testdata/bundle-with-package-manifest.cm.yaml",
			assertFunc: func(t *testing.T, manifestGot *Manifest) {
				assert.NotNil(t, manifestGot.Bundle)
				assert.NotNil(t, manifestGot.PackageManifest)

				assert.Equal(t, "etcd", manifestGot.PackageManifest.PackageName)
			},
		},

		{
			name:   "BundleWithMultiplePackageManifests",
			source: "testdata/bundle-with-multiple-package-manifests.cm.yaml",
			assertFunc: func(t *testing.T, manifestGot *Manifest) {
				assert.NotNil(t, manifestGot.Bundle)
				assert.NotNil(t, manifestGot.PackageManifest)

				name := manifestGot.PackageManifest.PackageName
				assert.True(t, name == "first" || name == "second")
			},
		},

		{
			name:   "BundleWithBuiltInKubeTypes",
			source: "testdata/bundle-with-kube-resources.cm.yaml",
			assertFunc: func(t *testing.T, manifestGot *Manifest) {
				assert.NotNil(t, manifestGot.Bundle)
				assert.NotNil(t, manifestGot.Bundle.Objects)

				objects := manifestGot.Bundle.Objects
				assert.Equal(t, 1, len(objects))
				assert.True(t, objects[0].GetKind() == "Foo")
			},
		},

		{
			name:   "BundleWithMultipleCsvs",
			source: "testdata/bundle-with-multiple-csvs.cm.yaml",
			assertFunc: func(t *testing.T, manifestGot *Manifest) {
				assert.NotNil(t, manifestGot.Bundle)

				csvGot, errGot := manifestGot.Bundle.ClusterServiceVersion()
				assert.NoError(t, errGot)
				assert.NotNil(t, csvGot)
				assert.True(t, csvGot.GetName() == "first" || csvGot.GetName() == "second")
			},
		},

		{
			name:   "BundleWithBadResource",
			source: "testdata/bundle-with-bad-resource.cm.yaml",
			assertFunc: func(t *testing.T, manifestGot *Manifest) {
				assert.NotNil(t, manifestGot.Bundle)

				csvGot, errGot := manifestGot.Bundle.ClusterServiceVersion()
				assert.NoError(t, errGot)
				assert.NotNil(t, csvGot)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cm := loadfromFile(t, tt.source)

			loader := NewBundleLoader()
			manifestGot, errGot := loader.Load(cm)

			assert.NoError(t, errGot)
			assert.NotNil(t, manifestGot)

			if tt.assertFunc != nil {
				tt.assertFunc(t, manifestGot)
			}
		})
	}
}

func loadfromFile(t *testing.T, path string) *corev1.ConfigMap {
	reader, err := os.Open(path)
	require.NoError(t, err, "unable to load from file %s", path)

	decoder := yaml.NewYAMLOrJSONDecoder(reader, 30)
	bundle := &corev1.ConfigMap{}
	err = decoder.Decode(bundle)
	require.NoError(t, err, "could not decode into configmap, file=%s", path)

	return bundle
}
