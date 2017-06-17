pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh "docker-compose build"
            }
        }
        stage('Test') {
            steps  {
                sh "docker-compose -f docker-compose.test.yml up --exit-code-from test"
            }
        }
        stage('Push') {
            steps {
                sh "docker-compose push"
            }
        }
        // stage('Deploy') {
        //     environment {
        //     }
        //     steps {

        //     }
        // }
    }
}