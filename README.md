# Registro

## Motivation

I find it annoying that I need a registry to do simple work when I'm using a simple standalone docker host. Even more when I'm in a test/dev machine and this process repeats forever.

#### The following sounds familiar (and annoying) to you?

```
# local machine

root@localhost:/hello$ docker build -t hello .
root@localhost:/hello$ docker tag hello some-registry.com/username/hello
root@localhost:/hello$ docker push some-registry.com/username/hello

# ssh into vps

root@vps:/$ docker pull some-registry.com/username/hello
root@vps:/$ docker tag some-registry.com/username/hello hello
root@vps:/$ docker run hello
Hello Registro!

```

#### Wouldn't it be awesome if we could just do this instead?

```
# local machine

root@localhost:/hello$ docker build -t hello .
root@localhost:/hello$ docker tag hello vps/hello
root@localhost:/hello$ docker push vps/hello

# ssh into vps

root@vps:/$ docker run hello
Hello Registro!

```

#### Or even

```
# local machine

root@localhost:/hello$ docker build -t hello 
root@localhost:/hello$ registro vps hello

# ssh into vps

root@vps:/$ docker run hello
Hello Registro!

```

## Where is the code?

This is an RFC, please let me know what you think in the issues first! ;)
