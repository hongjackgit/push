<?php

// namespace client
/*
* config example:{"vid":1,"content:":[{"target":1,"tname":"人群1","condition":[{"oper":"and","left":"1","right":"ccc"}]},{"target":1,"tname":"人群2","condition":[{"oper":"and","left":"1","right":"ccc"}]}]}
* oper: =,!=,in,not in,<,<=,>,>=
*/

// interface FileInterface {

	
// 	public function getFile();
// }

class Ab {

	private $aParams;
	private $oConfig;

	public function __construct($sBusinessName){
		$this->oConfig = new Config($sBusinessName);
	}

	public function allow($sBusinessName){
		
	}
}

class Oper{
	function in(){

	}
	function eq(){

	}
}

class Config{
	private $sBusinessName;
	private $sConfigKey;

	private $sConfigContent;

	// private $sPathPrefix = "/cronsun/etc/%s/current/";
	private $sPathPrefix = "./%s.json";

	public function __construct($sBusinessName){
		var_dump($sBusinessName);
		$this->sBusinessName = $sBusinessName;
	}

	public function getConfig(){
		$sPath = sprintf($this->sPathPrefix,$this->sBusinessName);
		$this->sConfigContent = file_get_contents($sPath);
		return $this->sConfigContent;
	}
}

$oConfig = new Config("task");

var_dump($oConfig->getConfig());
