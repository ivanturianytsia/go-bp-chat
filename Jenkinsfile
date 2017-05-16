node {
    checkout scm
    try {
      stage("Build") {
          sh 'bash build.sh alpine'
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
