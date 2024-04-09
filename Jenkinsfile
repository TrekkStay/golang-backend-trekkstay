pipeline {
    agent any

    parameters {
        choice(
            name: 'ACTION',
            choices: ['Removing container and image'],
            description: 'Select action to perform'
        )
    }

    stages {
        stage('Deploying application') {
            when {
                branch 'master'
            }

            steps {
                script {
                    // Remove existing container and image
                    sh 'rm -f trekkstay-backend || true'
                    sh 'docker rmi -f thanhanphan17/trekkstay-backend || true'

                    // Run the container
                    sh '''
                        docker container run \
                                --restart unless-stopped \
                                --env CONFIG_PATH=./env/prod.env \
                                --env MIGRATE=false \
                                --name trekkstay-backend \
                                --network trekkstay_network \
                                -dp 8888:8888 \
                                thanhanphan17/trekkstay-backend
                    '''
                }
            }
        }

        stage('Removing container and image') {
            when {
                expression { params.ACTION == 'Removing container and image' }
            }

            steps {
                script {
                    // Remove existing container and image
                    sh 'rm -f trekkstay-backend || true'
                    sh 'docker rmi -f thanhanphan17/trekkstay-backend || true'
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
