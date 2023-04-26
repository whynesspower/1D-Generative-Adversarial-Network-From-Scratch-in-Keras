/*
The back and forth processing required for GANs was time consuming - even in 2017 it
 could take days for a GAN to generate a realistic face. Another drawback of GANs is that they 
 are notoriously difficult and sensitive to train, suffering from failure to converge and mode collapse,
  where the generator starts preferring to generate only one or a few outputs that it knows have historically
  fooled the discriminator. GANs are still an active field of research and some recent developments such as 
  new cost functions, experience replay to reduce overfitting and mini-batch discrimination to address mode 
  collapse, have alleviated some of the difficulties of deploying them. Nowadays GANs have a wide range of 
  use cases including super resolution, face aging, background replacement, next frame prediction for videos and more.

*/