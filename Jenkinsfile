node {
    checkout scm

    stage("Build") {
        sh 'ls'
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
}
