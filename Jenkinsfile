podTemplate(
  label: 'pipeline',
  namespace: 'integration',
  serviceAccount: 'jnlp',
  containers: [
    containerTemplate(
      name: 'golang',
      image: 'golang:1.12.7',
      ttyEnabled: true
      ),
    containerTemplate(
      name: 'mysql',
      image: 'mysql:5.7.26',
      ttyEnabled: true,
      envVars: [
        envVar(key: "MYSQL_ROOT_PASSWORD", value: "password"),
        envVar(key: "MYSQL_DATABASE", value: "tabemap")
        ]
      ),
    containerTemplate(
      name: 'skaffold',
      image: 'gcr.io/k8s-skaffold/skaffold:latest',
      ttyEnabled: true,
      command: 'cat'
      )
  ],
  volumes: [
    hostPathVolume(hostPath: '/var/run/docker.sock', mountPath: '/var/run/docker.sock')
  ]
) {
  node ('pipeline') {
    withCredentials([
      usernamePassword(credentialsId: 'docker_id', usernameVariable: 'DOCKER_ID_USR', passwordVariable: 'DOCKER_ID_PSW'),
      usernamePassword(credentialsId: 'mysql_id', usernameVariable: 'DB_USER', passwordVariable: 'DB_PASSWORD')
    ]) {
      stage('Provisioning') {
        container('golang') {
          sh """
          """
        }
      }
      git 'https://github.com/ShotaKitazawa/tabemap-api'
      stage('Build') {
        container('golang') {
          sh """
            CGO_ENABLED=0 GOOS=linux go build -tags "mysql" -a -installsuffix cgo -o tabemap-api .
          """
        }
      }
      stage('Unit & Integration Test') {
        container('golang') {
          sh """
            DB_PASSWORD=password DB_NAME=tabemap go test -tags "mysql integration" -v -cover ./...
          """
        }
      }
      stage('Skaffold Run') {
        container('skaffold') {
          sh """
            docker login --username=$DOCKER_ID_USR --password=$DOCKER_ID_PSW
            perl -pi -e 's|^(  DB_USER: ).*\$|\$1'\$(echo -n $DB_USER | base64)'|g' kubernetes/manifest/manifest.yaml
            perl -pi -e 's|^(  DB_PASSWORD: ).*\$|\$1'\$(echo -n $DB_PASSWORD | base64)'|g' kubernetes/manifest/manifest.yaml
            skaffold run
          """
        }
      }
    }
  }
}

