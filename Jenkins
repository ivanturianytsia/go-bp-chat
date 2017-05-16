node {
    checkout scm

    stage("Build") {
        sh 'bash build.sh alpine'
        sh 'bash build.sh image'
    }
    stage("Publish") {
        withDockerRegistry([credentialsId: 'DockerHub']) {
            sh 'bash build.sh push'
        }
    }
}
