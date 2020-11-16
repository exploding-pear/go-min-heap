pipeline{
    agent any
    tools {
        go "go-1.15"
    }
    environment {
        GO111MODULE = 'on'
    }
    stages{
        stage('compile') {
            steps {
                sh """
                tar -C /usr/local -xzf go1.15.5.linux-arm64.tar.gz
                export PATH=$PATH:/usr/local/go/bin
                go version
                """
            }
        }
    }
}