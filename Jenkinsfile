pipeline{
    agent any
    tools {
        go 'go-1.15.4'
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