package deployment

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"

	"github.com/openshift/library-go/pkg/controller/factory"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resource/resourceapply"
	"github.com/openshift/library-go/pkg/operator/status"
	"github.com/openshift/library-go/pkg/operator/v1helpers"

	certmanoperatorinformers "github.com/openshift/cert-manager-operator/pkg/operator/informers/externalversions"
)

type CertManagerControllerSet struct {
	certManagerControllerStaticResourcesController factory.Controller
	certManagerControllerDeploymentController      factory.Controller
	certManagerWebhookStaticResourcesController    factory.Controller
	certManagerWebhookDeploymentController         factory.Controller
	certManagerCAInjectorStaticResourcesController factory.Controller
	certManagerCAInjectorDeploymentController      factory.Controller
}

func NewCertManagerControllerSet(
	kubeClient kubernetes.Interface,
	kubeInformersForNamespaces v1helpers.KubeInformersForNamespaces,
	kubeInformersForTargetNamespace informers.SharedInformerFactory,
	operatorClient v1helpers.OperatorClientWithFinalizers,
	certManagerOperatorInformers certmanoperatorinformers.SharedInformerFactory,
	kubeClientContainer *resourceapply.ClientHolder,
	eventRecorder events.Recorder,
	targetVersion string,
	versionRecorder status.VersionGetter,
	trustedCAConfigmapName string,
) *CertManagerControllerSet {
	return &CertManagerControllerSet{
		certManagerControllerStaticResourcesController: NewCertManagerControllerStaticResourcesController(operatorClient, kubeClientContainer, kubeInformersForNamespaces, eventRecorder),
		certManagerControllerDeploymentController:      NewCertManagerControllerDeploymentController(operatorClient, certManagerOperatorInformers, kubeClient, kubeInformersForTargetNamespace, eventRecorder, targetVersion, versionRecorder, trustedCAConfigmapName),
		certManagerWebhookStaticResourcesController:    NewCertManagerWebhookStaticResourcesController(operatorClient, kubeClientContainer, kubeInformersForNamespaces, eventRecorder),
		certManagerWebhookDeploymentController:         NewCertManagerWebhookDeploymentController(operatorClient, certManagerOperatorInformers, kubeClient, kubeInformersForTargetNamespace, eventRecorder, targetVersion, versionRecorder, trustedCAConfigmapName),
		certManagerCAInjectorStaticResourcesController: NewCertManagerCAInjectorStaticResourcesController(operatorClient, kubeClientContainer, kubeInformersForNamespaces, eventRecorder),
		certManagerCAInjectorDeploymentController:      NewCertManagerCAInjectorDeploymentController(operatorClient, certManagerOperatorInformers, kubeClient, kubeInformersForTargetNamespace, eventRecorder, targetVersion, versionRecorder, trustedCAConfigmapName),
	}
}

func (c *CertManagerControllerSet) ToArray() []factory.Controller {
	return []factory.Controller{
		c.certManagerControllerStaticResourcesController,
		c.certManagerControllerDeploymentController,
		c.certManagerWebhookStaticResourcesController,
		c.certManagerWebhookDeploymentController,
		c.certManagerCAInjectorStaticResourcesController,
		c.certManagerCAInjectorDeploymentController,
	}
}
