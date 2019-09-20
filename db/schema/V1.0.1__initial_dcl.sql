CREATE USER sillyhat_user identified BY 'sillyhat_user';
CREATE SCHEMA `sillyhat_user` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin ;
GRANT ALL ON sillyhat_user.* TO sillyhat_user;
CREATE USER sillyhat identified BY 'sillyhat';
GRANT ALL ON sillyhat_user.* TO sillyhat;