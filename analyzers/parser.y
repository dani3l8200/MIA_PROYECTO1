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
%token CONTINUE PAUSE 
//TOKENS FOR EXEC 
%token EXEC PATH  HYPHEN ARROW ROUTE
//TOKENS FOR MKDISK AND RMDISK 
%token  MKDISK SIZE UNIT NAME  NUMBERN K M ID STRTYPE RMDISK  
//TOKENS FOR FDISK 
%token FDISK ADD DELETE FIT TYPE B P E L BF FF WF FAST FULL 
//TOKENS FOR MOUNT 


//TOKEN PARA SEGUIR EN OTRA LINEA 
%type <token> CONTINUE PAUSE 
//TOKENS FOR EXEC 
%type <token>  EXEC PATH  HYPHEN ARROW ROUTE
//TOKENS FOR MKDISK AND RMDISK 
%type <token>   MKDISK SIZE UNIT NAME  NUMBERN K M ID STRTYPE RMDISK  
//TOKENS FOR FDISK 
%type <token>  FDISK ADD DELETE FIT TYPE B P E L BF FF WF FAST FULL 



//NO TERMINALES 
//OTHERS NO TERMINALS
%type <node> TYPE_NAMES
//NO TERMINALS PRINCIPALS 
%type <node> Input Command 
//NO TERMINALS FOR EXEC 
%type <node> Exec Paparams
//NO TERMINALS FOR MKDISK AND RMDIS
%type <node> Mkdisk Mkparams Rmdisk
//NO TERMINALS FOR FDISK 
%type <node> Fdisk FdiskParams


%%
//DEFENIS LA GRAMATICA
Input: /* empty */ { }
     | Command {$$ = $1; Root = $$}
     
Command: EXEC Exec {$$ = lwh.NodeF("EXEC","exec").Append($2)}
       | MKDISK Mkdisk {$$ = lwh.NodeF("MKDISK",$1).Append($2)}
       | RMDISK Rmdisk {$$ = lwh.NodeF("RMDISK",$1).Append($2)}
       | FDISK Fdisk {$$ = lwh.NodeF("FDISK",$1).Append($2)}
       | PAUSE {$$ = lwh.NodeF("PAUSE",$1)}
       | CONTINUE {$$ = lwh.NodeF("COMANDO CONTINUAR",$1)}
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





%% 
const src = `exec path`