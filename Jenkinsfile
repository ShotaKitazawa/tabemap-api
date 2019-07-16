def mysql_version = "5.7.26"
podTemplate(
  label: 'label',
  containers: [
    containerTemplate(name: 'golang', image: 'golang:1.12.7-alpine', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'mysql', image: 'mysql:5.7.26', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'postman', image: 'postman/newman:4.5.1-alpine', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'kaniko', image: 'gcr.io/kaniko-project/executor:latest', ttyEnabled: true, command: 'cat')
  ]
) {
  node ('label') {
    withCredentials([
      usernamePassword(credentialsId: 'docker_id', usernameVariable: 'DOCKER_ID_USR', passwordVariable: 'DOCKER_ID_PSW')
    ]) {
      stage('Info') {
        container('diuid') {
          sh """
            uname -a
            whoami
            pwd
            ls -al
          """
        }
      }
      git 'https://github.com/ShotaKitazawa/tabemap-api'
      stage('Unit Test') {
        container('golang') {
          sh """
            go test -v -cover ./...
          """
        }
      }
      stage('Build') {
        container('golang') {
          sh """
            go build -o tabemap-api main.go
          """
        }
      }
      stage('Run') {
        container('golang') {
          sh """
            ./tabemap-api #TODO
          """
        }
      }
      stage('Integration Test') {
        container('postman') {
          sh """
          """
        }
      }
      stage('Push') {
        container('diuid') {
          sh """
          """
        }
      }
    }
  }
}

