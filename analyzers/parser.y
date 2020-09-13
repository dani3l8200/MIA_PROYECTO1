%{
    // analyzers Contiene los analizadores usados para la lectura inicial
package analyzers


import "MIA-PROYECTO1/lwh"

var Root lwh.Node



%}

%union{
    node lwh.Node
    token string
}

/*-------------------------------TERMINALES----------------------------*/
//TOKEN PARA SEGUIR EN OTRA LINEA 
%token  PAUSE 
//TOKENS FOR EXEC 
%token EXEC PATH  HYPHEN ARROW ROUTE
//TOKENS FOR MKDISK AND RMDISK 
%token  MKDISK SIZE UNIT NAME  NUMBERN K M ID STRTYPE RMDISK  
//TOKENS FOR FDISK 
%token FDISK ADD DELETE FIT TYPE B P E L BF FF WF FAST FULL 
//TOKENS FOR MOUNT 
%token MOUNT
//TOKENS FOR UNMOUNT
%token UNMOUNT IDM
//TOKENS FOR SIMBOLS GENERALS
%token S_EQUAL COMMENT
//TOKENS FOR MKFS
%token MKFS IDN 
//TOKENS FOR MKFILE
%token MKFILE PCONT CONT 
//TOKENS FOR LOIN
%token LOGIN USER PWD 
//Tokens FOR LOGOUT
%token LOGOUT
//TOKENS FOR MKDIR 
%token MKDIR

%token REP RUTA 

//TOKEN PARA SEGUIR EN OTRA LINEA 
%type <token>  PAUSE 
//TOKENS FOR EXEC 
%type <token>  EXEC PATH  HYPHEN ARROW ROUTE
//TOKENS FOR MKDISK AND RMDISK 
%type <token>   MKDISK SIZE UNIT NAME  NUMBERN K M ID STRTYPE RMDISK  
//TOKENS FOR FDISK 
%type <token>  FDISK ADD DELETE FIT TYPE B P E L BF FF WF FAST FULL 
//TOKENS FOR MOUNT
%type <token> MOUNT 
//TOKENS FOR UNMOUNT
%type <token> UNMOUNT IDM
//TOKENS FOR SIMBOLS GENERALS 
%type <token> S_EQUAL COMMENT
//TOKENS FOR MKFS
%type <token> MKFS IDN
//TOKENS FOR MKFILE
%type <token> MKFILE PCONT CONT 
//TOKENS FOR LOGIN 
%type <token> LOGIN USER PWD 
//TOKENS FOR LOGOUT
%type <token> LOGOUT
//TOKENS FOR MKDIR
%type <token> MKDIR
//TOKENS FOR REPORTS 
%type <token> REP RUTA 

//NO TERMINALES 
//OTHERS NO TERMINALS
%type <node> TYPE_NAMES
//NO TERMINALS PRINCIPALS 
%type <node> Input Command ListCommands
//NO TERMINALS FOR EXEC 
%type <node> Exec Paparams
//NO TERMINALS FOR MKDISK AND RMDIS
%type <node> Mkdisk Mkparams Rmdisk
//NO TERMINALS FOR FDISK 
%type <node> Fdisk FdiskParams
//NO TERMINALS FOR MOUNT 
%type <node> Mount ParamsMount
//NO TERMINALS FOR UNMOUNT
%type <node> Unmount ListUnmount
//NO TERMINALS FOR MKFS
%type <node> Mkfs MkfsParams
//NO TERMINALS FOR MKFILE 
%type <node> Mkfile MkfileParams
//NO TERMINALS FOR LOGIN
%type <node> Login LoginParams
//NO TERMINALS FOR MKDIR
%type <node> Mkdir MkdirParams
//NO TERMINALS FOR REPORTS 
%type <node> Rep RepParams
%start Input

%%

Input: ListCommands {$$ = $1; Root = $$}
    ;
ListCommands: ListCommands Command {$$ = $1.Append($2)}
            | Command   {$$.Append($1)}
            ;
Command: EXEC Exec {$$ = lwh.NodeF("EXEC",$1).Append($2)}
       | MKDISK Mkdisk {$$ = lwh.NodeF("MKDISK",$1).Append($2)}
       | RMDISK Rmdisk {$$ = lwh.NodeF("RMDISK",$1).Append($2)}
       | FDISK Fdisk {$$ = lwh.NodeF("FDISK",$1).Append($2)}
       | PAUSE {$$ = lwh.NodeF("PAUSE",$1)}
       | UNMOUNT Unmount {$$ = lwh.NodeF("UNMOUNT",$1).Append($2)}
       | MOUNT Mount {$$ = lwh.NodeF("MOUNT",$1).Append($2)}
       | COMMENT {$$ = lwh.NodeF("COMENTARIO",$1)}
       | MKFS Mkfs {$$ = lwh.NodeF("MKFS",$1).Append($2)}
       | LOGIN Login {$$ = lwh.NodeF("LOGIN",$1).Append($2)}
       | LOGOUT {$$ = lwh.NodeF("LOGOUT",$1)}
       | MKFILE Mkfile {$$ = lwh.NodeF("MKFILE",$1).Append($2)}
       | MKDIR Mkdir {$$ = lwh.NodeF("MKDIR",$1).Append($2)}
       | REP Rep {$$ = lwh.NodeF("REP",$1).Append($2)}
       ;
Exec: Exec Paparams {$$ = $1.Append($2)}
    | Paparams {$$.Append($1)}
    ;
Mkdisk: Mkdisk Mkparams {$$ = $1.Append($2)}
      | Mkparams {$$.Append($1)}
      ;
Rmdisk: Rmdisk Paparams {$$ = $1.Append($2)}
      | Paparams {$$.Append($1)}
      ;
Fdisk: Fdisk FdiskParams {$$ = $1.Append($2)}
     | FdiskParams {$$.Append($1)}
     ;
Mount:  Mount ParamsMount {$$ = $1.Append($2)}
     |  ParamsMount {$$.Append($1)}
    ;
Unmount: Unmount ListUnmount {$$ = $1.Append($2)}
        | ListUnmount {$$.Append($1)}
       ;
Mkfs: Mkfs MkfsParams {$$ = $1.Append($2)}
    | MkfsParams {$$.Append($1)}
    ;
Login: Login LoginParams {$$ = $1.Append($2)}
    | LoginParams {$$.Append($1)}
    ;
Mkfile: Mkfile MkfileParams {$$ = $1.Append($2)}
      | MkfileParams {$$.Append($1)}
      ;
Mkdir: Mkdir MkdirParams {$$ = $1.Append($2)}
     | MkdirParams {$$.Append($1)}
     ;
Rep: Rep RepParams {$$ = $1.Append($2)}
    | RepParams {$$.Append($1)}
FdiskParams: Paparams {$$ = $1}
           | TYPE_NAMES {$$ = $1}
           | UNIT ARROW B {$$ = lwh.NodeF("UNIT",$3)}
           | TYPE ARROW P {$$ = lwh.NodeF("TYPE",$3)}
           | TYPE ARROW E {$$ = lwh.NodeF("TYPE",$3)}
           | TYPE ARROW L {$$ = lwh.NodeF("TYPE",$3)}
           | FIT ARROW BF {$$ = lwh.NodeF("FIT",$3)}
           | FIT ARROW FF {$$ = lwh.NodeF("FIT",$3)}
           | FIT ARROW WF {$$ = lwh.NodeF("FIT",$3)}
           | DELETE ARROW FAST {$$ = lwh.NodeF("DELETE",$3)}
           | DELETE ARROW FULL {$$ = lwh.NodeF("DELETE",$3)}
           | ADD ARROW NUMBERN {$$ = lwh.NodeF("ADD",$3)}
           | SIZE ARROW NUMBERN {$$ = lwh.NodeF("SIZE",$3)}
           | UNIT ARROW K  {$$ = lwh.NodeF("UNIT",$3)}
           | UNIT ARROW M  {$$ = lwh.NodeF("UNIT",$3)}
           ;

MkfsParams: IDN ARROW ID {$$ = lwh.NodeF("ID",$3)}
          | TYPE ARROW FAST {$$ = lwh.NodeF("TYPE",$3)}
          | TYPE ARROW FULL {$$ = lwh.NodeF("TYPE",$3)}
          | ADD ARROW NUMBERN {$$ = lwh.NodeF("ADD",$3)}
          | UNIT ARROW K {$$ = lwh.NodeF("UNIT",$3)}
          | UNIT ARROW M {$$ = lwh.NodeF("UNIT",$3)}
          | UNIT ARROW B {$$ = lwh.NodeF("UNIT",$3)}
          ;

LoginParams: USER ARROW ID {$$ = lwh.NodeF("USER",$3)}
           | PWD ARROW ID {$$ = lwh.NodeF("PWD",$3)}
           | PWD ARROW NUMBERN {$$ = lwh.NodeF("PWD",$3)}
           | IDN ARROW ID {$$ = lwh.NodeF("ID",$3)}
           ;
MkfileParams: Paparams {$$ = $1}
            | IDN ARROW ID {$$ = lwh.NodeF("ID",$3)}
            | PCONT {$$ = lwh.NodeF("P",$1)}
            | SIZE ARROW NUMBERN {$$ = lwh.NodeF("SIZE",$3)}
            | CONT ARROW ROUTE {$$ = lwh.NodeF("CONT",$3)}
            | CONT ARROW STRTYPE {$$ = lwh.NodeF("CONT",$3)}
            ;

MkdirParams: Paparams {$$ = $1}
           | IDN ARROW ID {$$ = lwh.NodeF("ID",$3)}
           | PCONT {$$ = lwh.NodeF("P",$1)}
           ;
RepParams: Paparams {$$ = $1}
         | IDN ARROW ID {$$ = lwh.NodeF("ID",$3)}
         | RUTA ARROW ROUTE {$$ = lwh.NodeF("RUTA",$3)}
         | RUTA ARROW STRTYPE {$$ = lwh.NodeF("RUTA",$3)}
         | NAME ARROW ID {$$ = lwh.NodeF("NAME",$3)}

ParamsMount: Paparams {$$ = $1}
           | TYPE_NAMES {$$ = $1}
           ;
ListUnmount: IDM ARROW ID {$$ = lwh.NodeF("ID",$3)}
        ;
Mkparams: Paparams {$$ = $1}
        | SIZE ARROW NUMBERN {$$ = lwh.NodeF("SIZE",$3)}
        | UNIT ARROW K  {$$ = lwh.NodeF("UNIT", $3)}
        | UNIT ARROW M  {$$ = lwh.NodeF("UNIT",$3)}
        | NAME ARROW ID {$$ = lwh.NodeF("NAME",$3)}
        ;
Paparams: PATH ARROW ROUTE {$$ = lwh.NodeF("PATH",$3)}
        | PATH ARROW STRTYPE {$$ = lwh.NodeF("PATH",$3)}
        ;
TYPE_NAMES: NAME ARROW ID {$$ = lwh.NodeF("NAME",$3)}
          | NAME ARROW STRTYPE {$$ = lwh.NodeF("NAME",$3)}
          ;
