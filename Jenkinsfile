#!/usr/bin/env groovy

pipeline {
    agent any
    tools {
        go 'Go'
    }
    environment{
        git_address = "https://github.com/chocoiii/WePanel.git"
        git_branch = "main"
    }

    stages {
        stage('1.拉取代码'){
            steps {
                dir('/home/WePanel'){
                    cleanWs()
                    git branch: "${git_branch}", url: "${git_address}"
                }
            }
        }


        stage('2.编译程序'){
            steps {
                dir('/home/WePanel'){
                    sh 'export GOROOT=/usr/local/go'
                    sh 'export PATH=\$PATH:\$GOROOT/bin'
                    sh 'go env -w GO111MODULE=on'
                    sh 'go env -w GOPROXY=https://goproxy.io,direct'
                    sh 'go env -w GOPATH=/root/go'
                    sh 'go build -o WePanel backend/main.go'
                }
            }
        }

        stage('3.部署程序'){
            steps {
                dir('/home/WePanel'){
                    script{
                        processId = sh(script: "lsof -t -i :5000 || true", returnStatus: true, returnStdout: true)
                        if (processId) {
                            echo "${processId}"
                            sh 'sudo kill -9 ${processId}'
                        }
                        else{
                            echo "None"
                        }
                        sh 'JENKINS_NODE_COOKIE=dontKillMe nohup /home/WePanel/WePanel >WePanel.log&'
                    }
                }

            }
        }
    }
}