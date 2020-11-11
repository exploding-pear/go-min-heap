pipeline{
    agent any
    tools {
        go "null"
    }
    environment {
        GO111MODULE = 'on'
    }
    stages{
        stage('compile') {
            steps {
                sh 'go build'
            }
        }
    }
}