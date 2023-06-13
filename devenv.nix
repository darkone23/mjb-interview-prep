{ pkgs, ... }:

{
  # https://devenv.sh/basics/
  # env.GREET = "devenv";

  # https://devenv.sh/packages/
  packages = [ 
    pkgs.git
    pkgs.just
    pkgs.sqlite
    pkgs.libargon2
    pkgs.nodePackages.typescript-language-server
    pkgs.nodePackages.vscode-css-languageserver-bin
    pkgs.nodePackages.vscode-html-languageserver-bin
    pkgs.nodePackages.vscode-json-languageserver-bin
  ];

  # https://devenv.sh/scripts/
  # scripts.hello.exec = "echo hello from $GREET";

  # enterShell = ''
  #   hello
  #   git --version
  # '';

  # https://devenv.sh/languages/
  languages.nix.enable = true;
  languages.go.enable = true;
  languages.javascript.enable = true;

  # https://devenv.sh/processes/
  processes.server.exec = "just server";

  # See full reference at https://devenv.sh/reference/options/
  difftastic.enable = true;

  # https://devenv.sh/pre-commit-hooks/
  pre-commit.hooks = {
    gotest.enable = true;
    gofmt.enable = true;
  };


}
