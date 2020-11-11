pipeline{
    agent any
    tools {
        go
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