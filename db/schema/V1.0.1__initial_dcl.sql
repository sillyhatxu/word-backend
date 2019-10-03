CREATE USER sillyhat_word identified BY 'sillyhat_word';
CREATE SCHEMA `sillyhat_word` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin ;
GRANT ALL ON sillyhat_word.* TO sillyhat_word;
GRANT ALL ON sillyhat_word.* TO sillyhat;