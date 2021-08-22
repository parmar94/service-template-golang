def ecrLoginPwd = ''

pipeline {   
  agent any
  environment {       
    registry = "717486009197.dkr.ecr.ap-south-1.amazonaws.com"       
    GOCACHE = "/tmp"
    SERVICE_NAME = "service-template-golang"   
  }   
  stages {       
    stage('Test & Build') { 
      steps {              
        // Create our project directory.
        sh 'docker run --rm -v "$PWD":/usr/src/github.com/Smart-Biz-Cloud-Solutions/${SERVICE_NAME} -w /usr/src/github.com/Smart-Biz-Cloud-Solutions/${SERVICE_NAME} golang:1.16 go build -v'       
        sh './${SERVICE_NAME}'
      }           
    }              
    stage('Build & Push Image') {
      when { 
        branch "v*.*"
      }
      steps {
        // build and publish release
        sh 'cd ${GOPATH}/src/github.com/Smart-Biz-Cloud-Solutions/${SERVICE_NAME}'
        sh "aws ecr get-login-password --region ap-south-1 | img login -u AWS --password-stdin  ${registry}" //-p ${ecrLoginPwd}
        sh "img build -t ${registry}/${SERVICE_NAME}:${GIT_LOCAL_BRANCH} ."  // when we run docker in this step, we're running it via a shell on the docker build-pod container, 
        sh "img push ${registry}/${SERVICE_NAME}:${GIT_LOCAL_BRANCH}" //$TAG_NAME"
      }
    }
  }
}
