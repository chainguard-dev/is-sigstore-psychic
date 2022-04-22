{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = { nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in rec {
        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            nixfmt
            go
            gocode
            gore
            gomodifytags
            gopls
            go-symbols
            gopkgs
            go-outline
            jq
          ];
        };
      });
}
