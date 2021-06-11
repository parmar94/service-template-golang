pipeline {   
  agent {
    kubernetes {
      label 'jenkins'  // all your pods will be named with this prefix, followed by a unique id
      idleMinutes 5 // how long the pod will live after no jobs have run on it
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
        container('docker') {  
          sh "curl -L http://www.google.com"
          sh "docker build -t ${SERVICE_NAME} ."  // when we run docker in this step, we're running it via a shell on the docker build-pod container, 
          sh "docker tag ${SERVICE_NAME}:latest ${registry}/${SERVICE_NAME}:$TAG_NAME"
          sh "docker image ls"
          // sh "docker push vividseats/promo-app:dev"        // which is just connecting to the host docker deaemon
        }
      }
    }
    // stage('Push Image') {
    //   steps {
    //     container('docker') {  
    //       sh "docker build -t vividseats/promo-app:dev ."  // when we run docker in this step, we're running it via a shell on the docker build-pod container, 
    //       sh "docker push vividseats/promo-app:dev"        // which is just connecting to the host docker deaemon
    //     }
    //   }
    // }
  }
}