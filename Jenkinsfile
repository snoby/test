// vi: ft=groovy
node {
  stage('pull code') {
    // pull the code to the workspace
    checkout scm
     echo "My branch is: ${env.BRANCH_NAME}"
     sh ' echo in sh: $BRANCH_NAME'
     echo "The Build num: ${env.BUILD_NUMBER}"
  }
  stage('prebuild') {
    echo "Running the prebuild stage"
  }
  stage('build') {
    echo "Running the build stage"
  }
  stage('test'){
    echo "Running the test stage"
  }
  stage('archive'){
    echo "Running the archive stage"
  }
  stage('deploy'){
    echo "Running the deploy stage"
    if (currentBuild.result == 'SUCCESS') { // <1>
      echo "Build status and test status is good deploying "
    }
    else {
      echo "Current build status is unstable not deploying"
    }
  }
  stage('cleanup'){
    echo "Running the cleanup stage"
  }

}
