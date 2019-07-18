podTemplate(
  label: 'pipeline',
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
  ]
) {
  node ('pipeline') {
    withCredentials([
      usernamePassword(credentialsId: 'docker_id', usernameVariable: 'DOCKER_ID_USR', passwordVariable: 'DOCKER_ID_PSW')
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
            CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tabemap-api .
          """
        }
      }
      stage('Unit & Integration Test') {
        container('golang') {
          sh """
            go test -v -cover ./...
          """
        }
      }
      stage('Skaffold Run') {
        container('skaffold') {
          sh """
            docker login --username=$DOCKER_ID_USR --password=$DOCKER_ID_PSW
            # TODO
            kubectl apply -f kubernetes/manifest/manifest.yaml
          """
        }
      }
    }
  }
}

