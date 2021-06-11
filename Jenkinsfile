def ecrLoginPwd = ''

pipeline {   
  agent {
    kubernetes {
      label 'jenkins'  // all your pods will be named with this prefix, followed by a unique id
      idleMinutes 30 // how long the pod will live after no jobs have run on it
      yamlFile 'jenkins-pod.yaml'  // path to the pod definition relative to the root of our project 
      defaultContainer 'golang'  // define a default container if more than a few stages use it, will default to jnlp container
    }
  }
  environment {       
    registry = "717486009197.dkr.ecr.ap-south-1.amazonaws.com"       
    GOCACHE = "/tmp"
    SERVICE_NAME = "service-template-golang"   
  }   
  stages {       
    stage('Build') {                 
      steps {               
        // Create our project directory.               
        sh 'cd ${GOPATH}/src'               
        sh 'mkdir -p ${GOPATH}/src/github.com/Smart-Biz-Cloud-Solutions/${SERVICE_NAME}'               
        // Copy all files in our Jenkins workspace to our project directory.                              
        sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/github.com/Smart-Biz-Cloud-Solutions/${SERVICE_NAME}'               
        // Build the app.               
        sh 'go build'                         
      }           
    }       
    stage('Test') {                    
      steps {                               
        sh 'go get github.com/cucumber/godog/cmd/godog'   
        sh 'godog'                      
      }       
    }       
    stage('Build Image') {
      // when {
      //   tag '*'
      // }
      steps {
        container('imagebuilder') {  
          sh "img build -t ${SERVICE_NAME} ."  // when we run docker in this step, we're running it via a shell on the docker build-pod container, 
          sh "img tag ${SERVICE_NAME}:latest ${registry}/${SERVICE_NAME}:latest" //$TAG_NAME"
          sh "img ls"
          // sh "docker push vividseats/promo-app:dev"        // which is just connecting to the host docker deaemon
        }
      }
    }
    stage('Push Image') {
      // when {
      //   tag '*'
      // }
      steps {
        withCredentials([aws(accessKeyVariable: 'AWS_ACCESS_KEY_ID', credentialsId: 'jenkins-ecr', secretKeyVariable: 'AWS_SECRET_ACCESS_KEY')]) {
          container('awscli') {
            sh "export AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}"
            sh "export AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}"
            sh "export AWS_DEFAULT_REGION=ap-south-1"
            sh "aws ecr get-login-password > pwd.txt"
            script {
              ecrLoginPwd = readFile('pwd.txt').trim()
            }
          }
        }

        container('imagebuilder') {  
          sh 'echo "${ecrLoginPwd}" | img login -u AWS --password-stdin ${registry}'
          sh "img push ${registry}/${SERVICE_NAME}:latest" //$TAG_NAME"
        }
      }
    }
  }
}