#!/usr/bin/env python3

import os, sys, glob

path = './lambdas'
buildFolderBase = '/tmp/taxifriend'
lambdaFolderBase = "taxifriend/lambdas"
files = os.listdir(path)

print('Building: taxifriend...')
for name in files:
    build = 'env GOOS=linux GOARCH=amd64 go build -o {0}/{1} {2}/{1}'.format(buildFolderBase, name, lambdaFolderBase)
    res = os.system(build)
    if res != 0:
      "Build Error"
      sys.exit(1)
      
print('Building: ok')

print('Package: processing')
packMap = {}
for name in files:
  
  compiledLambda = '{}/{}'.format(buildFolderBase, name)
  target = '{0}.zip'.format(compiledLambda)
  zipComd = 'zip -j {1} {0}'.format(compiledLambda, target)
  packMap[name] = target
  os.system(zipComd)
print('Package: ok')

print('Deploy: ...')
packageLambdas = glob.glob('{}/*.zip'.format(buildFolderBase))
for pack in packMap:
  funcUpdateCmdFormat = 'aws lambda update-function-code --function-name {0} --zip-file fileb://{1}'
  res = os.system(funcUpdateCmdFormat.format(pack, packMap[pack]))
  print('Deploy function: {0}, status: {1}'.format(pack, res))
