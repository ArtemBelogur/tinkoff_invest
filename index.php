<?php
$dbh = new PDO('mysql:dbname=tinkoff_invest;host=localhost', 'mysql', 'mysql');
$sth = $dbh->prepare("SELECT * FROM `obligations`");
$sth->execute();
$list = $sth->fetchAll(PDO::FETCH_ASSOC);
?>

<table>
    <thead>
        <tr>
            <th>ID</th>
            <th>time</th>
            <th>ofz_24021</th>
            <th>ofz_25084</th>
            <th>ofz_26207</th>
            <th>ofz_26211</th>
            <th>ofz_26212</th>
        </tr>
    </thead>
    <tbody>
        <?php foreach ($list as $row): ?>
        <tr>
            <td><?php echo $row['ID']; ?></td>
            <td><?php echo $row['time']; ?></td>
            <td><?php echo $row['ofz_24021']; ?> RUB</td>
            <td><?php echo $row['ofz_25084']; ?> RUB</td>
            <td><?php echo $row['ofz_26207']; ?> RUB</td>
            <td><?php echo $row['ofz_26211']; ?> RUB</td>
            <td><?php echo $row['ofz_26211']; ?> RUB</td>
        </tr>
        <?php endforeach; ?>
    </tbody>
</table>

<style type="text/css">
table {
    table-layout: fixed;
    padding: 0;
    width: 100%;
    border-collapse: collapse;
    border-spacing: 0;
    margin: 0 0 20px 0;
}

table th {
    vertical-align: middle;
    color: #777777;
    font-size: 10px;
    text-transform: uppercase;
    padding: 10px 5px;
    background: #dedede;
    text-align: center;
    border: 1px solid #c1c1c1;
}

table td {
    padding: 8px 5px;
    font-size: 12px;
    color: #000;
    border: 1px solid #e9e9e9;
    text-align: center;
    line-height: 16px;
    vertical-align: middle;
    font-size: 12px;
}
table tbody tr:nth-child(odd) td {
    background: #f4f4f4;
}

</style>