curl -sL https://deb.nodesource.com/setup_18.x | sudo -E bash -

sudo apt install nodejs

npx @backstage/create-app

cd k8sbackstage

yarn dev