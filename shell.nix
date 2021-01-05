{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.pgcli

    # keep this line if you use bash
    pkgs.bashInteractive
  ];
}
