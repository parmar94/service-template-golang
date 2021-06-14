def ecrLoginPwd = ''

pipeline {   
  agent {
    kubernetes {
      yamlFile 'jenkins-pod.yaml'  // path to the pod definition relative to the root of our project 
      defaultContainer 'go-img-builder'  // define a default container if more than a few stages use it, will default to jnlp container
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
    stage('Build & Push Image') {
      when { 
        branch "v*.*"
      }
      steps {
        sh 'cd ${GOPATH}/src/github.com/Smart-Biz-Cloud-Solutions/${SERVICE_NAME}'
        sh "aws ecr get-login-password --region ap-south-1 | img login -u AWS --password-stdin  ${registry}" //-p ${ecrLoginPwd}
        sh "img build -t ${registry}/${SERVICE_NAME}:${GIT_LOCAL_BRANCH} ."  // when we run docker in this step, we're running it via a shell on the docker build-pod container, 
        sh "img push ${registry}/${SERVICE_NAME}:${GIT_LOCAL_BRANCH}" //$TAG_NAME"
      }
    }
  }
}