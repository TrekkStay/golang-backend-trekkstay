pipeline {
    agent any

    parameters {
        choice(
            name: 'ACTION',
            choices: ['Removing all container and image'],
            description: 'Select action to perform'
        )
    }

    stages {
        stage('Packaging and Pushing Image') {
            when {
                branch 'master'
            }
            
            steps {
                script {
                    // Copy environment files to the workspace
                    sh 'sudo cp -r ~/env/ $WORKSPACE/'

                    // Build and push Docker image
                    docker.withRegistry('https://index.docker.io/v1/', 'dockerhub') {
                        sh 'docker build -t thanhanphan17/trekkstay-backend .'
                        sh 'docker push thanhanphan17/trekkstay-backend'
                    }
                }
            }
        }

        stage('Deploying application') {
            when {
                branch 'master'
            }

            steps {
                script {
                    // Remove existing container and image
                    sh 'docker rm -f trekkstay-backend || true'
                    sh 'docker rmi -f thanhanphan17/trekkstay-backend || true'

                    // Run the container
                    sh '''
                        docker container run \
                            --restart always \
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

        stage('Removing all container and image') {
            when {
                expression { params.ACTION == 'Removing all container and image' }
            }

            steps {
                script {
                    // Remove all containers and images
                    sh 'docker rm -f trekkstay-backend || true'
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
