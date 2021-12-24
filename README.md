# filter


kube-scheduler.yaml <sample>

    component: kube-scheduler
    tier: control-plane
  name: kube-scheduler
  namespace: kube-system
spec:
  containers:
  - command:
    - kube-scheduler
    - --authentication-kubeconfig=/etc/kubernetes/scheduler.conf
    - --authorization-kubeconfig=/etc/kubernetes/scheduler.conf
	  - --config=/etc/kubernetes/scheduler-config.yaml
	  - --bind-address=127.0.0.1
	  - --kubeconfig=/etc/kubernetes/scheduler.conf
	  - --leader-elect=false
	  - --port=0
	  #image: k8s.gcr.io/kube-scheduler:v1.19.16 //デフォルトスケジューラー
	  image: ryuto610/go-scheduler:0.2.1
	  imagePullPolicy: IfNotPresent
	  livenessProbe:
		  failureThreshold: 8
		  httpGet:
			  host: 127.0.0.1
			  path: /healthz
			  port: 10259
			  scheme: HTTPS
		  initialDelaySeconds: 10
		  periodSeconds: 10
		  timeoutSeconds: 15
	  name: kube-scheduler
	  resources:
		  requests:
			  cpu: 100m
	  startupProbe:
		  failureThreshold: 24
		  httpGet:
			  host: 127.0.0.1
			  path: /healthz
			  port: 10259
			  scheme: HTTPS
		  initialDelaySeconds: 10
		  periodSeconds: 10
		  timeoutSeconds: 15
	volumeMounts:
	- mountPath: /etc/kubernetes/scheduler.conf
    name: kubeconfig
    readOnly: true
	- mountPath: /etc/kubernetes/scheduler-config.yaml
	  name: scheduler-config
	  readOnly: true
  hostNetwork: true
  priorityClassName: system-node-critical
  volumes:
  - hostPath:
	    path: /etc/kubernetes/scheduler.conf
	    type: FileOrCreate
    name: kubeconfig
  - hostPath:
	    path: /etc/kubernetes/scheduler-config.yaml
	    type: File
    name: scheduler-config
status: {}
