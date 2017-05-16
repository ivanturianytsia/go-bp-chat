node {
    checkout scm
    try {
      stage("Build") {
          sh 'docker run --rm \
            -v "$PWD":/go/src/ivanturianytsia/goblueprints/chat \
            -w /go/src/ivanturianytsia/goblueprints/chat \
            blang/golang-alpine bash build.sh alpine'
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
    }
}
