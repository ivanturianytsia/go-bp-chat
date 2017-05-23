node {
    checkout scm
    try {
      stage("Build") {
          sh 'chmod -R 777 .'
          sh 'docker run --rm \
            -v /root/docker/jenkins/workspace/go-bp-chat:/go/src/chat \
            -w /go/src/chat \
            golang:latest bash -c "./build.sh alpine"'
          sh 'bash build.sh image'
      }
      stage("Publish") {
          withDockerRegistry([credentialsId: 'DockerHub']) {
              sh 'bash build.sh push'
          }
      }
      stage("Deploy") {
          withDockerRegistry([credentialsId: 'DockerHub']) {
            sh 'bash build.sh deploy'
          }
      }
    } catch(error) {
      throw error
    } finally {
      sh 'rm -rf *'
    }
}
