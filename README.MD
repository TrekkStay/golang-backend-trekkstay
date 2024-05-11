# Backend Build and Run Documentation


# 1. Install Docker
- You can find the installation documentation here: https://www.docker.com/

# 2. Install PostgresSQL and Redis
- The command to install and config Postgres and Redis is defined in Makefile
- You can run command (make sure you are in the root folder of backend repository)
  `
  make init_db
  `
  to install Postgres
  `
  make install_redis
  `
  to install redis

# 3. Config Enviroment
- In the folder `env` create 2 files `dev.env` and `prod.env` for development environment and production environment with the configuration below
```
BUILD_ENV=prod

VERSION=0.0.1

# SERVICE CONFIG
SERVICE_NAME=go-backend-trekkstay
SERVICE_HOST=localhost
SERVICE_PORT=8888
SERVICE_TIMEOUT=10

# DATABASE CONFIG
DB_HOST=
DB_PORT=5432
DB_NAME=trekkstay_db
DB_USERNAME=postgres
DB_PASSWORD=0000
DB_DRIVER=postgres

# Pool connection
MAX_OPEN_CONNECTIONS=100
MAX_IDLE_CONNECTIONS=25
# 5 minutes
CONNECTION_MAX_IDLE_TIME=300
# 10 minutes
CONNECTION_MAX_LIFE_TIME=600
# 30 minutes
CONNECTION_TIMEOUT=1800

# JWT
JWT_SECRET_KEY=
# 1 minute
VERIFY_TOKEN_EXPIRY=60
# 1 day
REFRESH_TOKEN_EXPIRY=604800
# 7 days
ACCESS_TOKEN_EXPIRY=86400

# MAIL CONFIG
MAIL_FROM=trekkstay.booking@gmail.com
MAIL_PASS= 
MAIL_SERVER=smtp.gmail.com
MAIL_PORT=587

# S3
S3_ACCESS_KEY=
S3_SECRET_KEY=
S3_REGION=ap-southeast-2
S3_BUCKET=trekkstay-storage
S3_CLOUDFRONT=

# Redis
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=""
REDIS_DB=0
```

# 4. Build and run backend
## 1. Option 1 - Pull from docker hub (already to use)
- My image at: https://hub.docker.com/r/thanhanphan17/trekkstay-backend/tags
- Run command
  `make run_image`
## 2. Option 2 - Build and run image from local
- Build and push image to Docker Hub
  `make push-iamge`
- Pull and run image from DockerHub
  `make run_image`

