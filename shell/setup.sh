# .bash_profile 설정
echo -e "\033[32;1m"🦆.bash_profile 설정"\033[m"
echo export GOROOT=/usr/local/go > ~/.bash_profile
echo export GOPATH="$HOME/go" >> ~/.bash_profile
echo source ~/.bashrc >> ~/.bash_profile
# oh-my-bash 설치
echo -e "\033[32;1m"🦆oh-my-bash 설치"\033[m"
bash -c "$(curl -fsSL https://raw.githubusercontent.com/ohmybash/oh-my-bash/master/tools/install.sh)"
# oh-my-bash 설정
echo -e "\033[32;1m"🦆oh-my-bash 설정"\033[m"
sed -i "s/OSH_THEME=\"font\"/OSH_THEME=\"powerline\"/g" ~/.bashrc
source ~/.bashrc
# jellybeans 설치
echo -e "\033[32;1m"🦆jellybeans 설치"\033[m"
mkdir -p ~/.vim/colors
cd ~/.vim/colors
curl -O https://raw.githubusercontent.com/nanotech/jellybeans.vim/master/colors/jellybeans.vim
echo colorscheme jellybeans >> ~/.vimrc
# vim-go 관련
echo -e "\033[32;1m"🦆vim-go 설치"\033[m"
# pathogen 설치
mkdir -p ~/.vim/autoload
cd ~/.vim/autoload
curl -LSso pathogen.vim https://tpo.pe/pathogen.vim
# .vimrc 설정
echo execute pathogen#infect() >> ~/.vimrc
echo syntax on >> ~/.vimrc
echo filetype plugin indent on >> ~/.vimrc
# vim-go 설치
mkdir -p ~/.vim/bundle
cd ~/.vim/bundle
git clone https://github.com/fatih/vim-go.git
