<?php
    include_once 'includes/dbh.inc.php'
?>

<!DOCTYPE html>
<html>
    <head>
        <title>RobinAI, A discord bot</title>
</head>
<body>

<p>I'm focusing on backend. To access your servers emoji stats, just find your server id below</p>

<?php

// set key to emoji_ID, value to array. Array will hold emoji_name, 1m, 3m, 6m, 12m
$tableList = array();
// This would be better if I turned it all into 1 run, where I take the DOB value, minus it off from currentDate to see how many days ago DOB was
$stmt = mysqli_stmt_init($conn);
$sql = "SHOW TABLES;";
$query=mysqli_query($conn, $sql);
if(mysqli_num_rows($query)>0) {
	while($row=mysqli_fetch_array($query)) {
		$tableList[] = $row[0];
	}
}

foreach ($tableList as $table) {
	$guildID = "";
	if (str_starts_with($table, "emojis_")) {
		$guildID = str_replace("emojis_", "", $table);
?>
		<a href="http://www.robinai.xyz/<?php echo $guildID; ?>">
    	<input type="submit" value='<?php echo "$guildID"; ?>' name="guildID"/>
    </a>

<?php
		echo $guildID;
		echo "\n";
		//$path = getcwd();
		mkdir ($_SERVER['DOCUMENT_ROOT']."/".$guildID, 0775);
		copy ("830676304487120956/index.php", $_SERVER['DOCUMENT_ROOT']."/".$guildID . "/index.php");
	}
}

?>
</body>
</html>
