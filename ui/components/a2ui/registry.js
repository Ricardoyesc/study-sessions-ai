import A2UIAudioPlayer from "./primitives/A2UIAudioPlayer.vue";
import A2UIButton from "./primitives/A2UIButton.vue";
import A2UICard from "./primitives/A2UICard.vue";
import A2UIColumn from "./primitives/A2UIColumn.vue";
import A2UIFlashcard from "./primitives/A2UIFlashcard.vue";
import A2UIImage from "./primitives/A2UIImage.vue";
import A2UIProgressBar from "./primitives/A2UIProgressBar.vue";
import A2UIQuizCard from "./primitives/A2UIQuizCard.vue";
import A2UIRichText from "./primitives/A2UIRichText.vue";
import A2UIRow from "./primitives/A2UIRow.vue";
import A2UISocraticDialog from "./primitives/A2UISocraticDialog.vue";
import A2UIText from "./primitives/A2UIText.vue";
import A2UIVideoPlayer from "./primitives/A2UIVideoPlayer.vue";

export const a2uiRegistry = {
  AudioPlayer: A2UIAudioPlayer,
  Button: A2UIButton,
  Card: A2UICard,
  Column: A2UIColumn,
  Flashcard: A2UIFlashcard,
  Image: A2UIImage,
  ProgressBar: A2UIProgressBar,
  QuizCard: A2UIQuizCard,
  RichText: A2UIRichText,
  Row: A2UIRow,
  SocraticDialog: A2UISocraticDialog,
  Text: A2UIText,
  VideoPlayer: A2UIVideoPlayer,
};

export function resolveA2UIComponent(type) {
  return a2uiRegistry[type] ?? null;
}
