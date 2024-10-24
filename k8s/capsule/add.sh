#!/bin/sh

if [ -z "$1" ]; then
    echo "Usage: $0 <username>"
    exit 1
fi

if [ -z $2 ]; then
    echo "Usage: $0 <username> <环境后缀:生产p,研发d>"
    exit 1
fi

USER=$1-$2


mkdir $USER
useradd $USER -s /bin/zsh

echo "User $USER added successfully"

cp -r ~/.ssh /home/$USER
cp -r ~/.oh-my-zsh /home/$USER
cp -r ~/.zshrc /home/$USER
cp -r ~/.vimrc /home/$USER
mkdir /home/$USER/.kube

# create-user.sh

bash ./update_user_kube_config.sh $USER $USER

# add text line 9
sed -i "9i\    namespace: $USER" $USER-$USER.kubeconfig

echo "User $USER added successfully"

cp $USER-$USER.* /home/$USER/.kube
mv $USER-$USER.* ./$USER

mv /home/$USER/.kube/$USER-$USER.kubeconfig /home/$USER/.kube/config

chown -R $USER:$USER /home/$USER

sed "s/<exp>/$USER/g" exp.yaml > ./$USER/$USER.yaml

kubectl apply -f ./$USER/$USER.yaml

echo "User $USER setup successfully"

su - $USER -c "kubectl create ns $USER"
