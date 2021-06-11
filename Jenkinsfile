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
    //registry = "magalixcorp/k8scicd"       
    GOCACHE = "/tmp"   
  }   
  stages {       
    stage('Build') {                 
      steps {               
        // Create our project directory.               
        sh 'cd ${GOPATH}/src'               
        sh 'mkdir -p ${GOPATH}/src/github.com/Smart-Biz-Cloud-Solutions/service-template-golang'               
        // Copy all files in our Jenkins workspace to our project directory.                              
        sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/github.com/Smart-Biz-Cloud-Solutions/service-template-golang'               
        // Build the app.               
        sh 'go build'                         
      }           
    }       
    stage('Test') {                    
      steps {                               
        // Create our project directory.               
        //sh 'cd ${GOPATH}/src'               
        //sh 'mkdir -p ${GOPATH}/src/github.com/Smart-Biz-Cloud-Solutions/service-template-golang'               
        // Copy all files in our Jenkins workspace to our project directory.                              
        //sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/github.com/Smart-Biz-Cloud-Solutions/service-template-golang'               
        // Remove cached test results.               
        //sh 'go clean -cache'               
        // Run Unit Tests. 
        //sh 'cd ${GOPATH}/src/github.com/Smart-Biz-Cloud-Solutions/service-template-golang'
        sh 'go get github.com/cucumber/godog/cmd/godog'   
        sh 'godog'                      
      }       
    }       
    // stage('Publish') {           
    //   environment {               
    //     registryCredential = 'dockerhub'           
    //   }           
    //   steps{               
    //     script {                   
    //       def appimage = docker.build registry + ":$BUILD_NUMBER"                   
    //       docker.withRegistry( '', registryCredential ) {                       
    //         appimage.push()                       
    //         appimage.push('latest')                   
    //       }               
    //     }           
    //   }       
    // }       
    // stage ('Deploy') {           
    //   steps {               
    //     script{                   
    //       def image_id = registry + ":$BUILD_NUMBER"                   
    //       sh "ansible-playbook  playbook.yml --extra-vars \"image_id=${image_id}\""               
    //     }           
    //   }       
    // }   
  }
}