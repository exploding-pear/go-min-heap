pipeline{
    agent any
    tools {
        go 'go-1.15'
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