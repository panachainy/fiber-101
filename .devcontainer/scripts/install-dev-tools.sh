# Install zsh
brew install zsh

# Install Oh my zsh 
/bin/bash -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# Select zsh
chsh -s $(which zsh)

# Verify shell
echo "$SHELL"

# Verify go
go version
