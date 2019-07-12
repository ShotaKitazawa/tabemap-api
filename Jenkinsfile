def golang_version = "1.12.7-alpine"
def mysql_version = "5.7.26"
podTemplate(
  label: 'diuid',
  containers: [
    containerTemplate(name: 'golang', image: 'golang:${golang_version}', ttyEnabled: true, command: 'cat')
    containerTemplate(name: 'diuid', image: 'weberlars/diuid:latest', ttyEnabled: true, command: 'cat')
    containerTemplate(name: 'postman', image: 'postman/newman:4.5.1-alpine', ttyEnabled: true, command: 'cat')
  ]
) {
  node {
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
        container('diuid') {
          sh """
            docker build . -t kanatakita/tabemap-api:$TODO
          """
        }
      }
      stage('Run') {
        container('diuid') {
          sh """
            docker run --rm -p 8080:8080 kanatakita/tabemap-api:$TODO
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
            docker tag 
            docker push 
          """
        }
      }
    }
  }
}

