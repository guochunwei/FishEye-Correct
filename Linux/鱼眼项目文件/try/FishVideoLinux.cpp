#include <iostream> // for standard I/O
#include <string>   // for strings
#include <iomanip>  // for controlling float print precision
#include <sstream>  // string to number conversion
#include <pthread.h>

#include "FishEyeCorrect.h"
#include  "opencv2/core/core.hpp"        // Basic OpenCV structures (cv::Mat, Scalar)
#include  "opencv2/imgproc/imgproc.hpp"  // Gaussian Blur
#include  "opencv2/highgui/highgui.hpp"// OpenCV window I/O

using namespace std;
using namespace cv;

int frameNumber = 0;
void  thrd_func1(Mat &orgImg,Mat &destImage);
void  thrd_func2(Mat &orgImg,Mat &destImage);
void  thrd_func3(Mat &orgImg,Mat &destImage);
CFishEyeCorrect fishCorrect;
	
int main(int argc, char *argv[])
{
//	string path = "C:\\Users\\Administrator\\Desktop\\360TSR_2.mp4";
	VideoCapture cap(argv[1]);
	frameNumber = cap.get(CV_CAP_PROP_FRAME_COUNT);
	cout<<"frameNumber "<<frameNumber<<endl;
	if(!cap.isOpened())
	{
		cout<<"读取视频失败！"<<endl;
		return 0;
	}
//	CFishEyeCorrect fishCorrect;
	int flag = 0;

	for(int i = 0;i < frameNumber;++i)
	{
		Mat frame;
	//	double posFrameIndex = cap.get(CV_CAP_PROP_POS_FRAMES);
		//pDlg->m_ctrlProgress.SetPos(posFrameIndex);
		//CString text;
		//text.Format("%d",(int)posFrameIndex);
		//pDlg->SetDlgItemText(IDC_TEXT,_T(text));

		cap >> frame;
		if (!frame.data)
			{
				break;
				}
		//clock_t start = clock();
	//	resize(frame,frame,Size(frame.cols/2,frame.rows/2));
		//clock_t start = clock();
 //	pDlg->ShowImage(frame,IDC_VIDEO1);
		//flip(frame,frame,-1);
		//Mat temp;
		//resize(frame,temp,Size(241,151));
		//pDlg->ShowImage(CSize(frame.cols,frame.rows),24,frame.data,1);



		 //
		if (flag == 0)
		{
			fishCorrect.GetArea(frame);
			flag =1;
		}
		Mat tempImg = frame(fishCorrect.correctArea);
		if (i ==0)
		{
			fishCorrect.CalculateParam(tempImg.cols,tempImg.rows);
			fishCorrect.FishImgIndex(tempImg);
			//FIC.Flag = 0;
		}
		//clock_t start = clock();
		// SM.Standard_circle(tempImg,imgCircle);

	/*	Mat dstImg1;
		Mat dstImg2;
		Mat dstImg3;

		fishCorrect.ImageCorrect(tempImg,dstImg1,dstImg2,dstImg3);*/
		//FIC.ImageCorrect(frame,dstImg1,dstImg2,dstImg3);
		/*clock_t end = clock();
		CString time;
		time.Format("%d\n%d",end - middle,middle - start);
		pDlg->MessageBox(time);*/
		// writer << dstImg1;


		//double angle = -90;
		//double scale = 1;
		//Mat WarpImg1;
		//Mat WarpImg2;
		//Mat WarpImg3;
		////pDlg->RotationwarpAffine(dstImg1,WarpImg1,angle,scale);
		//pDlg->RotationwarpAffine(dstImg2,WarpImg2,angle,scale);
		//pDlg->RotationwarpAffine(dstImg3,WarpImg3,angle,scale);

		//pDlg->ShowImage(CSize(WarpImg1.cols,WarpImg1.rows),24,WarpImg1.data,2);
		//pDlg->ShowImage(CSize(dstImg2.cols,dstImg2.rows),24,dstImg2.data,3);
		//pDlg->ShowImage(CSize(WarpImg2.cols,WarpImg2.rows),24,WarpImg2.data,4);
		
		
		 pthread_t tid1,tid2,tid3;
     void *tret;
     // 创建线程tid1，线程函数thrd_func1
     if (pthread_create(&tid1,NULL,thrd_func1,tempImg)!=0) {
         printf("Create thread 1 error!\n");
         exit(1);
     }
     // 创建线程tid2,线程函数thrd_func2
     if (pthread_create(&tid2,NULL,thrd_func2,tempImg)!=0) {
         printf("Create thread 2 error!\n");
         exit(1);
     }
			// 创建线程tid2,线程函数thrd_func2
     if (pthread_create(&tid3,NULL,thrd_func3,tempImg)!=0) {
         printf("Create thread 3 error!\n");
         exit(1);
     }

     // 等待线程tid1结束，线程函数返回值放在tret中
     if (pthread_join(tid1,&tret)!=0){
         printf("Join thread 1 error!\n");
         exit(1);
     }
 
    // printf("Thread 1 exit code: %d.\n",(int *)tret);
     // 等待tid2结束，线程函数返回值放在tret中
    if (pthread_join(tid2,&tret)!=0){
         printf("Join thread 2 error!\n");
        exit(1);
     }
		 // 等待tid2结束，线程函数返回值放在tret中
     if (pthread_join(tid3,&tret)!=0){
         printf("Join thread 3 error!\n");
         exit(1);
     }	

	/*	namedWindow("Video0");
		namedWindow("Video1");
		namedWindow("Video2");
		namedWindow("Video3");
		imshow("Video0",tempImg);
		imshow("Video1",dstImg1);
		imshow("Video2",dstImg2);
		imshow("Video3",dstImg3);*/
		waitKey(1);

	//	clock_t end = clock();
		/*CString time;
		time.Format("%d",end - start);
		pDlg->MessageBox(time);*/
		//imwrite("C:\\Users\\Administrator\\Desktop\\001.jpg",imgCircle);
		// if(waitKey(10) >= 0) break;
	}
	cout<<"处理完毕！"<<endl;
	return 0;
}

void  thrd_func1(Mat &orgImg)
{
	Mat &destImage;
	fishCorrect.ImageCorrect(orgImg,destImage,fishCorrect.ImgIndex1);
	namedWindow("Video1");
	imshow("Video1",destImage);
	pthread_exit(NULL); 
	}
	
void  thrd_func2(Mat &orgImg)
{
	Mat &destImage;
	fishCorrect.ImageCorrect(orgImg,destImage,fishCorrect.ImgIndex2);
	namedWindow("Video2");
	imshow("Video2",destImage);
	pthread_exit(NULL); 
	
	}

void  thrd_func3(Mat &orgImg)
{
	Mat &destImage;
	fishCorrect.ImageCorrect(orgImg,destImage,fishCorrect.ImgIndex3);
	namedWindow("Video3");
	imshow("Video3",destImage);
	pthread_exit(NULL); 
	}