#!/usr/bin/env groovy

pipeline {
    agent any
    tools {
        go 'Go'
        nodejs 'NodeJs'
    }
    environment{
        git_address = "https://github.com/chocoiii/WePanel.git"
        git_branch = "main"
    }

    stages {
        stage('测试'){
            steps {
                echo env.CHANGE_TARGET
            }
        }
        stage('1.拉取代码'){
            when {
                expression {
                    return env.CHANGE_ACTION == 'closed'
                }
            }
            steps {
                dir('/home/WePanel'){
                    cleanWs()
                    git branch: "${git_branch}", url: "${git_address}"
                }
            }
        }

        stage('2.后端编译'){
            when {
                expression {
                    return env.CHANGE_ACTION == 'closed'
                }
            }
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

        stage('3.后端部署'){
            when {
                expression {
                    return env.CHANGE_ACTION == 'closed'
                }
            }
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

        stage('4.前端构建'){
            when {
                expression {
                    return env.CHANGE_ACTION == 'closed'
                }
            }
            steps {
                dir('/home/WePanel/frontend/webPanel'){
                    script{
                        sh 'npm install --ignore-scripts --registry=https://mirrors.cloud.tencent.com/npm/'
                        sh 'npm run build'
                    }
                }

            }
        }

        stage('5.前端部署'){
            when {
                expression {
                    return env.CHANGE_ACTION == 'closed'
                }
            }
            steps {
                dir('/home/WePanel/frontend/webPanel'){
                    script{
                        processId = sh(script: "lsof -t -i :5173 || true", returnStatus: true, returnStdout: true)
                        if (processId) {
                            echo "${processId}"
                            sh 'sudo kill -9 ${processId}'
                        }
                        else{
                            echo "None"
                        }
                        sh 'JENKINS_NODE_COOKIE=dontKillMe nohup npm run dev >webPanel.log&'
                    }
                }

            }
        }
    }
}