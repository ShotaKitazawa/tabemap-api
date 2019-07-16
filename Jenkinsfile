def mysql_version = "5.7.26"
podTemplate(
  label: 'label',
  containers: [
    containerTemplate(name: 'golang', image: 'golang:1.12.7-alpine', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'diuid', image: 'weberlars/diuid:latest', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'postman', image: 'postman/newman:4.5.1-alpine', ttyEnabled: true, command: 'cat')
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
            go get -u github.com/golang/dep/cmd/dep
            dep ensure
            go test -v -cover ./...
          """
        }
      }
      stage('Build') {
        container('diuid') {
          sh """
            docker build . -t kanatakita/tabemap-api:latest
          """
        }
      }
      stage('Run') {
        container('diuid') {
          sh """
            docker pull mysql:${mysql_version}
            docker run --rm -p 8080:8080 kanatakita/tabemap-api:latest
            docker run --rm -p 3306:3306 mysql:${mysql_version}
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

