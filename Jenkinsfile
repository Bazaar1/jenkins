pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                cd apiserver/ &&
                go build
            }
        }
    }
}
