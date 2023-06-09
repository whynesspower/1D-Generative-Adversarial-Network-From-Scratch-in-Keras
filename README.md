# 1D-Generative-Adversarial-Network-From-Scratch-in-Keras

GAN or Generative Adversarial Network are a type of Deep Learning architecture for training powerful generator models.

Its is a backbone to almost every genrative model out there may it be OpenAI's DALL-E 2, Google Brain's Imagen, Midjourney or Stable Diffusion. Yann Lecun, Chief AI Scientist at Meta, said it was “the most interesting idea in the last 10 years in machine learning”
A one dimentinal simple implementation of GAN using TensorFlow's library Keras

GAN is composed of a generator model and a descriminator model competing with each other. Hence the name 'adversarial'. 

This project is based on GAN, Sequential model from Keras. We will have one node for input, 25 nodes in hidden layer and one output layer for our discriminator model. 

*Checkout resouces section for more details and some references to learn more about GANs models*

### Dependancies

1. pip install graphviz
2. pip install pydot
3. pip install tensorflow
4. pip install keras
5. pip install numpy
6. pip install matplotlib

`if you are facing any dependancy error its because you don't have installed some library which above libraries internally depend on, read the error message for such library and update`

We are doing total of 10,000 epochs and plotting a progress graph every 2,000 epochs. Five differnt plots. Below I am pasting the first and last one of those.

### This below first plot is created after 2,000 iterations and shows real (red) vs. fake (blue) samples. The model performs poorly initially with a cluster of generated points only in the positive input domain, although with the right functional relationship

![image](https://user-images.githubusercontent.com/77494053/234606186-294dd4f2-e8e4-4583-ac38-9e0688ecb678.png)

### This below second shows real (red) vs. fake (blue) after 10,000 epochs

![image](https://user-images.githubusercontent.com/77494053/234606488-0977be36-8151-4b25-82a4-5bc003ec877c.png)

## My Learnings-

1. How to develop separate discriminator and generator models, as well as a composite model for training the generator via the discriminator’s predictive behavior.

2. How to subjectively evaluate generated samples in the context of real examples from the problem domain.
