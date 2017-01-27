// vi: ft=groovy
node {
  // pull the code to the workspace
  checkout scm


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
