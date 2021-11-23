<?php
    include_once $_SERVER['DOCUMENT_ROOT'] . '/includes/dbh.inc.php'
?>

<!DOCTYPE html>
<html>
    <head>
        <title></title>
</head>
<body>
<style>
table {
  font-family: arial, sans-serif;
  border-collapse: collapse;
  width: 100%;
}

td, th {
  border: 1px solid #dddddd;
  text-align: left;
  padding: 8px;
}

tr:nth-child(even) {
  background-color: #dddddd;
}
</style>


<?php
// Stores the guilds emojis as emojilist[id]->name
$emojilist = array();

// set key to emoji_ID, value to array. Array will hold emoji_name, 1m, 3m, 6m, 12m
$data = array();
$date = strtotime(date("Y-m-d"));
// This would be better if I turned it all into 1 run, where I take the DOB value, minus it off from currentDate to see how many days ago DOB was
$stmt = mysqli_stmt_init($conn);
$guildID = basename(getCwd());

$sql = "SELECT * from emojiGuild_" . $guildID . ";";
$query=mysqli_query($conn, $sql);
if(mysqli_num_rows($query)>0) {
	while($row=mysqli_fetch_object($query)) {
        $emojilist[$row->emoji_ID] = $row->emoji_name;
    }
}
$sql = "SELECT * from emojis_" . $guildID . ";";
$query=mysqli_query($conn, $sql);
if(mysqli_num_rows($query)>0):
	while($row=mysqli_fetch_object($query)) {
		// hopefully this is php speak for checking the date difference
		// update: 
		$dateDOB = strtotime($row->DOB);
		$days = ($dateDOB - $date) / 86400;
		// intialize the key for the associative multi dimensional array
		if (!array_key_exists($row->emoji_ID, $data)) {
			$data[$row->emoji_ID] = array($row->emoji_name, 0, 0, 0, 0);
		}
		// increment the array values
		if ($days < 365){
			$data[$row->emoji_ID][4]++; // 12m
		}
		if ($days < 183) {
			$data[$row->emoji_ID][3]++; // 6m
		}
		if ($days < 92) {
			$data[$row->emoji_ID][2]++; // 3m
		}
		if ($days < 30) {
			$data[$row->emoji_ID][1]++; // 1m
		}
	}
?>



<table>
    <tr>
        <th align="center"><form><input type=submit name="btn_submit" value="Emoji ID" style="width:100%"></form></td>
        <th align="center"><form><input type=submit name="btn_submit" value="Emoji Name" style="width:100%"></form></td>
        <th align="center"><form><input type=submit name="btn_submit" value="1 Month" style="width:100%"></form></td>
	<th align="center"><form><input type=submit name="btn_submit" value="3 Month" style="width:100%"></form></td>
	<th align="center"><form><input type=submit name="btn_submit" value="6 Month" style="width:100%"></form></td>
	<th align="center"><form><input type=submit name="btn_submit" value="12 Month" style="width:100%"></form></td>
    </tr>
    <?php
    function sort1($a, $b) {
        if ($a[1]==$b[1]) return 0;
        if ($a[1] < $b[1]) return 1;
        else return -1;
    }
    function sort2($a, $b) {
        if ($a[2]==$b[2]) return 0;
        if ($a[2] < $b[2]) return 1;
        else return -1;
    }
    function sort3($a, $b) {
        if ($a[3]==$b[3]) return 0;
        if ($a[3] < $b[3]) return 1;
        else return -1;
    }
    function sort4($a, $b) {
        if ($a[4]==$b[4]) return 0;
        if ($a[4] < $b[4]) return 1;
        else return -1;
    }
    switch ($_REQUEST['btn_submit']) {
        case "Emoji ID":
    		krsort($data);
    		break;
    	case "Emoji Name":
    		arsort($data);
    		break;
    	case "1 Month":
    		uasort($data, 'sort1');
    		break;
    	case "3 Month":
    		uasort($data, 'sort2');
    		break;
    	case "6 Month":
    		uasort($data, 'sort3');
    		break;
    	case "12 Month":
    		uasort($data, 'sort4');
		break;
    }
    // looping
    foreach($data as $key => $val): ?>
    <tr>
        <td align="center"><?php echo $key;  // Emoji id ?></td>
        <td align="center"><?php echo $val[0]; // Emoji name ?></td>
        <td align="center"><?php echo $val[1]; // 1m  ?></td>
	<td align="center"><?php echo $val[2]; // 3m  ?></td>
	<td align="center"><?php echo $val[3]; // 6m  ?></td>
	<td align="center"><?php echo $val[4]; // 12m  ?></td>
    </tr>
    <?php endforeach; ?>
</table>
<?php
// no result show
else: ?>
<h3>No Results found.</h3>
<?php endif;

// Remove the used emojis from emojilist, so the unused emojis are left in emojilist.
foreach ($emojilist as $key => $value) {
    if (array_key_exists($key, $data)) {
        unset($emojilist[$key]);
    }
}
?>
<table>
    <tr>
        <th align="center">Emoji ID</td>
        <th align="center">Emoji Name</td>
    </tr>
    <?php if (!empty($emojilist)):
        foreach($emojilist as $emoji_ID => $emoji_name): ?>
        <tr>
        <td align="center"><?php echo $emoji_ID; ?></td>
        <td align="center"><?php echo $emoji_name; ?></td>
    </tr>
    <?php endforeach; ?>
    <?php endif; ?>
</table>
</body>
</html>
