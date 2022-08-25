SOURCE_IMAGE = 'europe-west4-docker.pkg.dev/cr-tap/vscode/vmiauw'
LOCAL_PATH = '.'
NAMESPACE = 'dev'

k8s_custom_deploy(
    'vmiauw',
    apply_cmd="tanzu apps workload apply -f workload.yaml --live-update" +
        " --local-path " + LOCAL_PATH +
        " --source-image " + SOURCE_IMAGE +
        " --namespace " + NAMESPACE +
        " --yes >/dev/null" +
        " && kubectl get workload vmiauw --namespace " + NAMESPACE + " -o yaml",
    delete_cmd="tanzu apps workload delete -f workload.yaml --namespace " + NAMESPACE + " --yes" ,
    deps=['pom.xml', './target/classes'],
    container_selector='workload',
    live_update=[
        sync('./target/classes', '/workspace/BOOT-INF/classes')
    ]
)

k8s_resource('vmiauw', port_forwards=["8080:8080"],
    extra_pod_selectors=[{'serving.knative.dev/service': 'vmiauw'}])