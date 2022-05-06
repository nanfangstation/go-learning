package main

import (
	"time"

	"code.byted.org/infcp/vke-resource/pkg/clustermanager"
	"code.byted.org/infcp/vke-resource/pkg/util/misc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

"apiVersion: install.istio.io/v1alpha1\nkind: IstioOperator\nmetadata:\n  name: installed-state\n  namespace: istio-system\nspec:\n  components:\n    pilot:\n      k8s:\n        env:\n          - name: PILOT_ENABLE_MYSQL_FILTER\n            value: \"true\"\n          - name: ISTIO_META_DNS_CAPTURE\n            value: \"true\"\n  profile: demo\n  meshConfig:\n    outboundTrafficPolicy:\n      mode: REGISTRY_ONLY\n  tag: 1.9.1"

var _ = Describe("helm interface", func() {
	var (
		err     error
		cOpts   *clustermanager.ClusterOptions
		cls     clustermanager.ClusterWarp
		helmCli *Client
		as      AddonService
		opt     *AddonOptions
		name    string
		ns      string
	)

	BeforeEach(func(done Done) {
		var ID string
		var kubeconfig []byte

		ID, kubeconfig, err = GenerateFakeCluster(cfg)
		Expect(err).NotTo(HaveOccurred())

		cOpts = &clustermanager.ClusterOptions{
			ClusterConfig: clustermanager.NewOptions().ClusterConfig,
			ID:            ID,
			Kubeconfig:    kubeconfig,
		}

		cls, err = clustermanager.NewClusterWarp(cOpts)
		Expect(err).NotTo(HaveOccurred())

		helmCli = &Client{
			ClientGetter: cls.RESTClientGetter(),
			Client:       cls.GetAPIReader(),
			Logger:       logf.Log,
		}

		as = NewAddonService()

		name = "nginx"
		ns = "default"
		opt = &AddonOptions{
			Name:            name,
			Namespace:       ns,
			ChartVersion:    "1.0.0",
			CreateNamespace: true,
			Install:         true,
			MaxHistory:      5,
			Timeout:         time.Minute * 1,
		}
		close(done)
	}, timeoutSeconds)

	AfterEach(func(done Done) {
		_ = as.UninstallRelease(helmCli, opt)
		close(done)
	}, timeoutSeconds)

	It("InstallRelease", func() {
		rls, err := as.InstallRelease(helmCli, opt)
		Expect(err).NotTo(HaveOccurred())
		Expect(rls).ToNot(BeNil())

		items, err := as.ListRelease(helmCli, &FilterOption{})
		Expect(err).NotTo(HaveOccurred())
		Expect(len(items)).To(Equal(1))

		items, err = as.ListRelease(helmCli, &FilterOption{
			Namespace: &ns,
			Filter:    &name,
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(len(items)).To(Equal(1))
	})

	It("InstallRelease have ChartName", func() {
		opt.ChartName = name
		rls, err := as.InstallRelease(helmCli, opt)
		Expect(err).NotTo(HaveOccurred())
		Expect(rls).ToNot(BeNil())
	})

	It("GetRelease", func() {
		rls, err := as.InstallRelease(helmCli, opt)
		Expect(err).NotTo(HaveOccurred())
		Expect(rls).ToNot(BeNil())

		rls, err = as.GetRelease(helmCli, opt)
		Expect(err).NotTo(HaveOccurred())
		Expect(rls.Name).To(Equal(name))
	})

	It("ApplyRelease", func() {
		rls, err := as.ApplyRelease(helmCli, opt)
		Expect(err).NotTo(HaveOccurred())
		Expect(rls).ToNot(BeNil())

		rls, err = as.ApplyRelease(helmCli, opt)
		Expect(err).NotTo(HaveOccurred())
		Expect(rls.Name).To(Equal(name))

		OverrideValues := "replicaCount: 12"
		opt.OverrideValues = misc.String2bytes(OverrideValues)
		rls, err = as.ApplyRelease(helmCli, opt)
		Expect(err).NotTo(HaveOccurred())
		Expect(rls.Name).To(Equal(name))
	})

	It("ApplyRelease with crds", func() {
		name = "security-scan"
		ns = "kube-system"
		withCRDs := &AddonOptions{
			Name:            name,
			Namespace:       ns,
			ChartVersion:    "v1.0.4",
			CreateNamespace: true,
			Install:         true,
			MaxHistory:      5,
			Timeout:         time.Minute * 1,
		}

		rls, err := as.ApplyRelease(helmCli, withCRDs)
		Expect(err).NotTo(HaveOccurred())
		Expect(rls).ToNot(BeNil())
		Expect(rls.Name).To(Equal(name))

		err = as.UninstallRelease(helmCli, withCRDs)
		Expect(err).NotTo(HaveOccurred())
	})
})
