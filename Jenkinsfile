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
                sh "docker-compose -f docker-compose.test.yml up --build --exit-code-from test"
            }
        }
        stage('Push') {
            steps {
                sh "docker-compose push"
            }
        }
        stage('Deploy') {
            environment {
                DOCKER_HOST = "${env.SWARM_HOST}"
                DOCKER_CERT_PATH = "${env.SWARM_CERT_PATH}"
                DOCKER_TLS_VERIFY = "1"
            }
            steps {
                sh "docker stack deploy -c docker-compose.production.yml vote"
            }
        }
    }
}