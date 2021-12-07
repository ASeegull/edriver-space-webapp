<?php 
 session_start(); 
 if($_SESSION['token'] !== 'vhid-is-ok'){ 
 
 if($_SESSION['send'] == 'no'){$ip = 'В панель управління Borders( http://adv-webstudio.esy.es ) було здійснено вхід з ip: ' . $_SERVER["REMOTE_ADDR"] .
  '. Перевірити місцезнаходження: http://xseo.in/ipinfo .'; echo 'ok';
mail("skskuzan@gmail.com", "Вхід в панель управління", $ip, "From: agnertech@gmail.com \r\n");
$_SESSION['send'] = 'yes';}
 $_SESSION['token'] = 'vhid-is-ne-ok';
 header ('Location: ./index.php');
 }

	?>