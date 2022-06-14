package goTest

import (
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"net/http"
	"testing"
)

//文件测试

func TestFile(t *testing.T) {
	var path = "C:\\Users\\daida\\Desktop\\bili\\8f18a4bc-57f6-4c2a-8b21-fafbb0973bc2.mp4"
	file, err := gfile.Open(path)
	if err != nil {
		var errstring any = "读取错误"
		panic(errstring)
	}
	fmt.Println(file.Name())
	bytes := gfile.GetBytesByTwoOffsetsByPath(path, 0, 10)
	fmt.Println(bytes)
}

//path为本地文件路劲
func play(path string, httpreq *http.Request, httpres *http.Response) {
	//file, err := gfile.Open("")
	//
	//RandomAccessFile targetFile = null;
	//OutputStream outputStream = null;
	//try {
	//outputStream = response.getOutputStream();
	//response.reset();
	////获取请求头中Range的值
	//String rangeString = request.getHeader("Range");
	//
	////打开文件
	//File file = new File(path);
	//if (file.exists()) {
	////使用RandomAccessFile读取文件
	//targetFile = new RandomAccessFile(file, "r");
	//long fileLength = targetFile.length();
	//long requestSize = (int) fileLength;
	////分段下载视频
	//if (StringUtils.hasText(rangeString)) {
	////从Range中提取需要获取数据的开始和结束位置
	//long requestStart = 0, requestEnd = 0;
	//String[] ranges = rangeString.split("=");
	//if (ranges.length > 1) {
	//String[] rangeDatas = ranges[1].split("-");
	//requestStart = Integer.parseInt(rangeDatas[0]);
	//if (rangeDatas.length > 1) {
	//requestEnd = Integer.parseInt(rangeDatas[1]);
	//}
	//}
	//if (requestEnd != 0 && requestEnd > requestStart) {
	//requestSize = requestEnd - requestStart + 1;
	//}
	////根据协议设置请求头
	//response.setHeader(HttpHeaders.ACCEPT_RANGES, "bytes");
	//response.setHeader(HttpHeaders.CONTENT_TYPE, "video/mp4");
	//if (!StringUtils.hasText(rangeString)) {
	//response.setHeader(HttpHeaders.CONTENT_LENGTH, fileLength + "");
	//} else {
	//long length;
	//if (requestEnd > 0) {
	//length = requestEnd - requestStart + 1;
	//response.setHeader(HttpHeaders.CONTENT_LENGTH, "" + length);
	//response.setHeader(HttpHeaders.CONTENT_RANGE, "bytes " + requestStart + "-" + requestEnd + "/" + fileLength);
	//} else {
	//length = fileLength - requestStart;
	//response.setHeader(HttpHeaders.CONTENT_LENGTH, "" + length);
	//response.setHeader(HttpHeaders.CONTENT_RANGE, "bytes " + requestStart + "-" + (fileLength - 1) + "/"
	//+ fileLength);
	//}
	//}
	////文段下载视频返回206
	//response.setStatus(HttpServletResponse.SC_PARTIAL_CONTENT);
	////设置targetFile，从自定义位置开始读取数据
	//targetFile.seek(requestStart);
	//} else {
	////如果Range为空则下载整个视频
	//response.setHeader(HttpHeaders.CONTENT_DISPOSITION, "attachment; filename=test.mp4");
	////设置文件长度
	//response.setHeader(HttpHeaders.CONTENT_LENGTH, String.valueOf(fileLength));
	//}
	//
	////从磁盘读取数据流返回
	//byte[] cache = new byte[4096];
	//try {
	//while (requestSize > 0) {
	//int len = targetFile.read(cache);
	//if (requestSize < cache.length) {
	//outputStream.write(cache, 0, (int) requestSize);
	//} else {
	//outputStream.write(cache, 0, len);
	//if (len < cache.length) {
	//break;
	//}
	//}
	//requestSize -= cache.length;
	//}
	//} catch (IOException e) {
	//// tomcat原话。写操作IO异常几乎总是由于客户端主动关闭连接导致，所以直接吃掉异常打日志
	////比如使用video播放视频时经常会发送Range为0- 的范围只是为了获取视频大小，之后就中断连接了
	//log.info(e.getMessage());
	//}
	//} else {
	//throw new RuntimeException("文件路劲有误");
	//}
	//outputStream.flush();
	//} catch (Exception e) {
	//log.error("文件传输错误", e);
	//throw new RuntimeException("文件传输错误");
	//}finally {
	//if(outputStream != null){
	//try {
	//outputStream.close();
	//} catch (IOException e) {
	//log.error("流释放错误", e);
	//}
	//}
	//if(targetFile != null){
	//try {
	//targetFile.close();
	//} catch (IOException e) {
	//log.error("文件流释放错误", e);
	//}
	//}
	//}
}
