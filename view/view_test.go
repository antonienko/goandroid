package view

import (
	"encoding/xml"
	"testing"
)

func TestGetByMatchingResourceAfterText(t *testing.T) {
	xmlHierarchy := new(Hierarchy)
	err := xml.Unmarshal([]byte(rawxml2), xmlHierarchy)
	if err != nil {
		t.Error("Failed to load xml", err.Error())
	}
	views, err := xmlHierarchy.ConvertToViews()
	if err != nil {
		t.Error(err.Error())
	}
	viewJavi, _ := views.GetByMatchingResourceAfterText("contactpicker_call_button", "Javi")
	callButtonIdForJavi := 43
	if viewJavi.TreeIndex != callButtonIdForJavi {
		t.Errorf("Found the wrong node. Found %d and expected %d", viewJavi.TreeIndex, callButtonIdForJavi)
	}
	viewAntonio, _ := views.GetByMatchingResourceAfterText("contactpicker_call_button", "Antonio")
	callButtonIdForAntonio := 31
	if viewAntonio.TreeIndex != callButtonIdForAntonio {
		t.Errorf("Found the wrong node. Found %d and expected %d", viewAntonio.TreeIndex, callButtonIdForAntonio)
	}
}

var rawxml2 = `<?xml version='1.0' encoding='UTF-8' standalone='yes' ?>
<hierarchy rotation="0">
	<node index="1" text="0" resource-id="" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][320,480]">
		<node index="2" text="1" resource-id="" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][320,480]">
			<node index="3" text="2" resource-id="" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,24][320,480]">
				<node index="4" text="3" resource-id="com.whatsapp:id/action_bar_root" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,24][320,480]">
					<node index="5" text="4" resource-id="android:id/content" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,24][320,480]">
						<node index="6" text="5" resource-id="com.whatsapp:id/root_layout" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,24][320,480]">
							<node index="7" text="6" resource-id="" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,24][320,480]">
								<node index="8" text="7" resource-id="" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,24][320,80]">
									<node index="9" text="8" resource-id="com.whatsapp:id/toolbar" class="android.view.ViewGroup" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,24][320,80]">
										<node index="10" text="9" resource-id="" class="android.widget.ImageButton" package="com.whatsapp" content-desc="Navigate up" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,24][56,80]" />
										<node index="11" text="9" resource-id="" class="android.widget.TextView" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[72,29][136,57]" />
										<node index="12" text="9" resource-id="" class="android.widget.TextView" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[72,57][132,75]" />
										<node index="13" text="9" resource-id="" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[144,32][184,72]">
											<node index="14" text="13" resource-id="com.whatsapp:id/progress_bar" class="android.widget.ProgressBar" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[152,40][176,64]" />
										</node>
										<node index="15" text="9" resource-id="" class="android.support.v7.widget.al" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[184,24][320,80]">
											<node index="16" text="15" resource-id="com.whatsapp:id/menuitem_search" class="android.widget.TextView" package="com.whatsapp" content-desc="Search" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="true" password="false" selected="false" bounds="[184,28][232,76]" />
											<node index="17" text="15" resource-id="com.whatsapp:id/menuitem_new_contact" class="android.widget.TextView" package="com.whatsapp" content-desc="New contact" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="true" password="false" selected="false" bounds="[232,28][280,76]" />
											<node index="18" text="15" resource-id="" class="android.widget.ImageView" package="com.whatsapp" content-desc="More options" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="true" password="false" selected="false" bounds="[280,28][320,76]" />
										</node>
									</node>
								</node>
								<node index="19" text="7" resource-id="android:id/content" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,80][320,480]">
									<node index="20" text="19" resource-id="android:id/list" class="android.widget.ListView" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="true" scrollable="false" long-clickable="true" password="false" selected="false" bounds="[0,80][320,480]">
										<node index="21" text="20" resource-id="" class="android.widget.RelativeLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,80][300,156]">
											<node index="22" text="21" resource-id="" class="android.widget.RelativeLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,80][300,156]">
												<node index="23" text="22" resource-id="com.whatsapp:id/contact_selector" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="true" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,80][82,156]">
													<node index="24" text="23" resource-id="com.whatsapp:id/contactpicker_row_photo" class="android.widget.ImageView" package="com.whatsapp" content-desc="Anto" checkable="false" checked="false" clickable="true" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[14,91][68,145]" />
													<node index="25" text="23" resource-id="com.whatsapp:id/selection_check" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[50,124][74,148]" />
												</node>
												<node index="26" text="22" resource-id="" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,106][194,129]">
													<node index="27" text="26" resource-id="" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,106][180,129]">
														<node index="28" text="27" resource-id="" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,106][180,129]">
															<node index="29" text="Anto" resource-id="com.whatsapp:id/contactpicker_row_name" class="android.widget.TextView" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,106][118,129]" />
														</node>
													</node>
												</node>
												<node index="30" text="22" resource-id="com.whatsapp:id/buttons" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[194,80][300,156]">
													<node index="31" text="30" resource-id="com.whatsapp:id/contactpicker_call_button" class="android.widget.ImageButton" package="com.whatsapp" content-desc="Voice call" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[194,80][244,156]" />
													<node index="32" text="30" resource-id="com.whatsapp:id/contactpicker_videocall_button" class="android.widget.ImageButton" package="com.whatsapp" content-desc="Video call" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[244,80][294,156]" />
												</node>
											</node>
										</node>
										<node index="33" text="20" resource-id="" class="android.widget.RelativeLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,157][300,233]">
											<node index="34" text="33" resource-id="" class="android.widget.RelativeLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,157][300,233]">
												<node index="35" text="34" resource-id="com.whatsapp:id/contact_selector" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="true" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,157][82,233]">
													<node index="36" text="35" resource-id="com.whatsapp:id/contactpicker_row_photo" class="android.widget.ImageView" package="com.whatsapp" content-desc="Javi" checkable="false" checked="false" clickable="true" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[14,168][68,222]" />
													<node index="37" text="35" resource-id="com.whatsapp:id/selection_check" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[50,201][74,225]" />
												</node>
												<node index="38" text="34" resource-id="" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,183][194,206]">
													<node index="39" text="38" resource-id="" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,183][180,206]">
														<node index="40" text="39" resource-id="" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,183][180,206]">
															<node index="41" text="Javi" resource-id="com.whatsapp:id/contactpicker_row_name" class="android.widget.TextView" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,183][112,206]" />
														</node>
													</node>
												</node>
												<node index="42" text="34" resource-id="com.whatsapp:id/buttons" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[194,157][300,233]">
													<node index="43" text="42" resource-id="com.whatsapp:id/contactpicker_call_button" class="android.widget.ImageButton" package="com.whatsapp" content-desc="Voice call" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[194,157][244,233]" />
													<node index="44" text="42" resource-id="com.whatsapp:id/contactpicker_videocall_button" class="android.widget.ImageButton" package="com.whatsapp" content-desc="Video call" checkable="false" checked="false" clickable="true" enabled="true" focusable="true" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[244,157][294,233]" />
												</node>
											</node>
										</node>
										<node index="45" text="20" resource-id="" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,234][300,296]">
											<node index="46" text="45" resource-id="" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,234][300,296]">
												<node index="47" text="46" resource-id="com.whatsapp:id/image" class="android.widget.ImageView" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[14,238][68,292]" />
												<node index="48" text="46" resource-id="" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,234][300,296]">
													<node index="49" text="48" resource-id="" class="android.widget.TextView" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,234][185,296]" />
												</node>
											</node>
										</node>
										<node index="50" text="20" resource-id="" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,297][300,359]">
											<node index="51" text="50" resource-id="" class="android.widget.LinearLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,297][300,359]">
												<node index="52" text="51" resource-id="com.whatsapp:id/image" class="android.widget.ImageView" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[14,301][68,355]" />
												<node index="53" text="51" resource-id="" class="android.widget.FrameLayout" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,297][300,359]">
													<node index="54" text="53" resource-id="" class="android.widget.TextView" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[82,297][192,359]" />
												</node>
											</node>
										</node>
									</node>
								</node>
							</node>
						</node>
					</node>
				</node>
			</node>
		</node>
		<node index="55" text="1" resource-id="android:id/statusBarBackground" class="android.view.View" package="com.whatsapp" content-desc="" checkable="false" checked="false" clickable="false" enabled="true" focusable="false" focused="false" scrollable="false" long-clickable="false" password="false" selected="false" bounds="[0,0][320,24]" />
	</node>
</hierarchy>`
