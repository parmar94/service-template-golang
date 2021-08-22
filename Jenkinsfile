pipeline {   
  agent any
  environment {       
    registry = "717486009197.dkr.ecr.ap-south-1.amazonaws.com"       
    GOCACHE = "/tmp"
    SERVICE_NAME = "service-template-golang"   
  }   
  stages {       
    stage('Build') { 
      steps {              
        sh 'docker run --rm -v "$PWD":/usr/src/github.com/Smart-Biz-Cloud-Solutions/${SERVICE_NAME} -w /usr/src/github.com/Smart-Biz-Cloud-Solutions/${SERVICE_NAME} golang:1.16 go build -v'
      }           
    }              
    stage('Build & Push Image') {
      //when { 
      //  branch "v*.*"
      //}
      steps {
        // build and publish release
        sh "password=\$(docker run --rm -i amazon/aws-cli ecr get-login-password --region ap-south-1)"
        sh "docker login -u AWS -p $password   ${registry}"
        sh 'docker build -t ${registry}/${SERVICE_NAME}:${GIT_LOCAL_BRANCH} .'
        sh "docker push ${registry}/${SERVICE_NAME}:${GIT_LOCAL_BRANCH}" //$TAG_NAME"
      }
    }
  }
}
